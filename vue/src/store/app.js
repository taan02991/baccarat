module.exports = {
  types: [
    // this line is used by starport scaffolding
    { type: "game", fields: ["state", "participant", "result", "resultHash", "bet", ] },
    { type: "game/start", fields: ["id"] },
    { type: "user", fields: ["name", "addr"] }
  ]
};
