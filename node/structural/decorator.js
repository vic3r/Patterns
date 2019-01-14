const carModel = function(options = {}) {
  return {
    getCars: function() {},
    getCarById: function() {},
    limit: options.limit || 10
  }
}

let car = carModel();
console.log(car.limit);

let fancyCar = (carModel) => {
  carModel.limit += 50;
  carModel.addStuff = () => {};
  carModel.removeStuff = () => {};

  return carModel;
}

console.log(fancyCar(car).limit);
