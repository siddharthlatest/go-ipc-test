const _ = require('lodash');
a = () => {
  console.log("hello world sum(2 + 2): ", 2+2)
  console.log("ext lib example: ", _.chunk(['a', 'b', 'c', 'd'], 2))
}
a();
