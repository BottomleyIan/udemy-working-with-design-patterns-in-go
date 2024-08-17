document.addEventListener("DOMContentLoaded", function(evt) {
  fetch("api/dog-breeds", { method: "GET" })
    .then((response) => response.json())
    .then((response) => {
      if (!response.length) {
        return;
      }
      new window.simpleDatatables(".dog-breeds", {
        perPage: 25,
        columns: [
          {
            select: 1,
            render: function(data, td, rowIndex) {
              return `<a href="/dog-breeds/${response[rowIndex].id}">${data[0].data}</a>`;
            },
          },
          {
            select: 4,
            render: function(data) {
              return `<div class="text-center">${data[0].data}</div>`;
            },
          },
          {
            select: 5,
            render: function(data) {
              return `<div class="text-center">${data[0].data}</div>`;
            },
          },
          { select: [0, 2, 3, 6, 7, 9], hidden: true },
        ],
      });
    })
    .catch((error) => {
      console.error("Error:", error);
    });
});
