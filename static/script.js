const app = document.getElementsByClassName("app")[0];

const template = () => {
  const supreme = "supreme";
  return `
          <h1>Merhaba</h1>
          <p>Merhaba Dünya</p>
          <p>${supreme}</p>
        `;
};

const style = {
  border: "1px solid red",
  color: "blue",
};

app.setAttribute("style", `${style}`);
app.innerHTML = template();
