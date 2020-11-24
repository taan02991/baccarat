module.exports = {
  types: [
    // this line is used by starport scaffolding
<<<<<<< HEAD
		{ type: "game", fields: ["roomNo", "round", "participant", "result", "resultHash", ] },
=======
    { type: "game", fields: ["state", "participant", "result", "resultHash", "bet", ] },
    { type: "game/start", fields: ["id"] },
    { type: "game/bet", fields: ["id", "side", "amount"] },
    { type: "game/participant", fields: ["id", "action"] },
    { type: "user", fields: ["name", "addr"] }
>>>>>>> bcf50511433ee7e1c80c53bd95832d37d7057b95
  ]
};
