module.exports = {
  types: [
    // this line is used by starport scaffolding
		{ type: "game", fields: ["state", "participant", "result", "resultHash", "bet", ] },
    { type: "user", fields: ["name", "addr"] }
  ]
};
