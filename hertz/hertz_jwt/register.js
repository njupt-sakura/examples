new Promise((resolve, reject) => {
  fetch("http://127.0.0.1:8888/register", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      username: "whoami",
      email: "161043261@qq.com",
      password: "pass",
    }),
  })
    .then((res) => res.json())
    .then((res) => {
      console.log("res:", res);
      resolve(res);
    })
    .catch((err) => {
      console.error("err:", err);
      reject(err);
    });
});
