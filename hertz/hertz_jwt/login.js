new Promise((resolve, reject) => {
  fetch("http://127.0.0.1:8888/login", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      account: "whoami",
      password: "pass",
    }),
  })
    .then((res) => res.json())
    .then((res) => {
      console.log(res);
      resolve(res);
    })
    .catch((err) => {
      console.error(err);
      reject(err);
    });
});

// curl http://127.0.0.1:8888/auth/ping?token=<token>
