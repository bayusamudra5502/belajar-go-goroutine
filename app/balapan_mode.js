// Race condition di JS
// Tidak akan terjadi karena sistem pemrosesan async itu sendiri

async function belalang() {
  let x = 0

  for (let i = 0; i < 1000; i++) {
    (function () {
      return new Promise((resolve) => {
        console.log(i)

        for (let j = 0; j < 100; j++) {
          x++
        }

        resolve()
      })
    })()
  }

  return x
}

belalang().then(console.log)

console.log("Halo, Dunia")
