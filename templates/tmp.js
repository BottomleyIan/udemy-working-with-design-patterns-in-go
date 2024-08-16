let dogFactoryBtn = document.getElementById("dog-factory-btn");
let catFactoryBtn = document.getElementById("cat-factory-btn");

let dogFactoryOutput = document.getElementById("dog-factory-output");

document.addEventListener("DOMContentLoaded", function() {
  dogFactoryBtn.addEventListener("click", function() {
    fetch("api/dog-from-factory", { method: "GET" })
      .then((response) => response.json())
      .then((data) => {
        if (data.error) {
          dogFactoryOutput.innerHTML = data.error;
          return;
        }
        dogFactoryOutput.innerHTML = JSON.stringify(data, null, 2);
      })
      .catch((error) => {
        console.error("Error:", error);
      });
  });

  catFactoryBtn.addEventListener("click", function() {
    fetch("api/cat-from-factory", { method: "GET" })
      .then((response) => response.json())
      .then((data) => {
        if (data.error) {
          dogFactoryOutput.innerHTML = data.error;
          return;
        }
        dogFactoryOutput.innerHTML = JSON.stringify(data, null, 2);
      })
      .catch((error) => {
        console.error("Error:", error);
      });
  });
});
