const KEY = Symbol.for("My.App.Namespace.Foo");

let globalSymbols = Object.getOwnPropertySymbols(global);
let hasKey = (globalSymbols.indexOf(KEY) > -1);

if (!hasKey) {
  global[KEY] = {
    "config": "value"
  };
}

let singleton = {};

Object.defineProperty(singleton, "instance", {
  get: function () {
    return global[KEY];
  }
});

Object.freeze(singleton);

module.exports = singleton;
