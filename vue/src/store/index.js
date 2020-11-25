import Vue from "vue";
import Vuex from "vuex";
import axios from "axios";
import app from "./app.js";
import * as bip39 from "bip39";
import {
  Secp256k1Wallet,
  SigningCosmosClient,
  makeCosmoshubPath
} from "@cosmjs/launchpad";

Vue.use(Vuex);

const API = process.env.VUE_APP_API;
const ADDRESS_PREFIX = "cosmos";

export default new Vuex.Store({
  state: {
    app,
    name: "",
    account: {},
    chain_id: "",
    data: {},
    client: null
  },
  mutations: {
    accountUpdate(state, { account }) {
      state.account = account;
    },
    chainIdSet(state, { chain_id }) {
      state.chain_id = chain_id;
    },
    entitySet(state, { type, body }) {
      const updated = {};
      updated[type] = body;
      state.data = { ...state.data, ...updated };
    },
    clientUpdate(state, { client }) {
      state.client = client;
    },
    setName(state, name) {
      state.name = name;
    }
  },
  actions: {
    async init({ dispatch }) {
      await dispatch("chainIdFetch");
      dispatch("entityFetch", { type: "game" });
      dispatch("entityFetch", { type: "user" });
    },
    async chainIdFetch({ commit }) {
      const node_info = (await axios.get(`/node_info`)).data.node_info;
      commit("chainIdSet", { chain_id: node_info.network });
    },
    async accountSignIn({ commit }, { mnemonic }) {
      return new Promise((resolve, reject) => {
        Secp256k1Wallet.fromMnemonic(
          mnemonic,
          makeCosmoshubPath(0),
          ADDRESS_PREFIX
        ).then(wallet => {
          wallet.getAccounts().then(([{ address }]) => {
            const url = `/auth/accounts/${address}`;
            axios.get(url).then(res => {
              const acc = res.data;
              if (acc.result.value.address === address) {
                const account = acc.result.value;
                const client = new SigningCosmosClient(API, address, wallet);
                commit("accountUpdate", { account });
                commit("clientUpdate", { client });
                axios.get(`/baccarat/user/${address}`).then(res => {
                  commit("setName", res.data.result.name);
                });
                resolve(account);
              } else {
                reject("Account doesn't exist.");
              }
            });
          });
        });
      });
    },
    async accountRegister({ state }, name) {
      const mnemonic = bip39.generateMnemonic();
      const wallet = await Secp256k1Wallet.fromMnemonic(
        mnemonic,
        makeCosmoshubPath(0),
        ADDRESS_PREFIX
      );
      const [{ address }] = await wallet.getAccounts();
      const body = {
        name: name,
        addr: address
      };
      const fundingWallet = await Secp256k1Wallet.fromMnemonic(
        process.env.VUE_APP_FUNDING_ACCOUNT,
        makeCosmoshubPath(0),
        ADDRESS_PREFIX
      );
      const fundingAddr = (await fundingWallet.getAccounts())[0].address;
      const client = new SigningCosmosClient(API, fundingAddr, fundingWallet);
      const { chain_id } = state;
      const creator = client.senderAddress;
      const base_req = { chain_id, from: creator };
      const req = { base_req, creator, ...body };
      const { data } = await axios.post(`/${chain_id}/user`, req);
      const { msg, fee, memo } = data.value;
      await client.signAndPost(msg, fee, memo);
      return mnemonic;
    },
    async entityFetch({ state, commit }, { type }) {
      const { chain_id } = state;
      const url = `/${chain_id}/${type}`;
      const body = (await axios.get(url)).data.result;
      commit("entitySet", { type, body });
    },
    async accountUpdate({ state, commit }) {
      const url = `/auth/accounts/${state.client.senderAddress}`;
      const acc = (await axios.get(url)).data;
      const account = acc.result.value;
      commit("accountUpdate", { account });
    },
    async entitySubmit({ state }, { type, body }) {
      const { chain_id } = state;
      const creator = state.client.senderAddress;
      const base_req = { chain_id, from: creator };
      const req = { base_req, creator, ...body };
      const { data } = await axios.post(`/${chain_id}/${type}`, req);
      const { msg, fee, memo } = data.value;
      await state.client.signAndPost(msg, fee, memo);
      return data.value.msg[0];
    }
  }
});
