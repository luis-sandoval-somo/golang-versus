let limit = 1000000;
function test100000() {
  for (let counter = 0; counter <= limit; counter++) {
    // if (counter % 100000 == 0) {
    console.log(counter);
    // }
    // if (counter == limit) {
    //   console.log("done", counter);
    // }
  }
}
test100000();
