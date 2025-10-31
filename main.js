import https from "https";

const TodoList = [];

function getUrl(url) {
  return new Promise((resolve, reject) => {
    https
      .get(url, (res) => {
        let data = "";

        res.on("data", (chunk) => {
          data += chunk;
        });

        res.on("end", () => {
          resolve(data);
        });
      })
      .on("error", (err) => {
        reject(err);
      });
  });
}

async function getTodo(id) {
  return getUrl(`https://jsonplaceholder.typicode.com/todos/${id}`);
}

async function main() {
  const start = performance.now();

  const promises = [];
  for (let i = 1; i <= 100; i++) {
    promises.push(getTodo(i));
  }

  const results = await Promise.all(promises);
  TodoList.push(...results);

  const end = performance.now();

  console.log("✅ Done. Total Todos:", TodoList.length);
  console.log("⏱️ Time:", (end - start).toFixed(2), "ms");
}

main().catch(console.error);
