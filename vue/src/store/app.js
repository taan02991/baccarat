module.exports = {
  types: [
    // this line is used by starport scaffolding
    {
      type: "game",
      fields: ["state", "participant", "result", "resultHash", "bet"]
    },
    { type: "game/start", fields: ["id"] },
    { type: "game/bet", fields: ["id", "side", "amount"] },
    { type: "game/participant", fields: ["id", "action"] },
    { type: "user", fields: ["name", "addr"] }
  ]
};
