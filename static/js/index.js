function CreateMessage(data) {
  const message = document.createElement("div");
  message.classList.add("message");

  const header = document.createElement("div")
  header.classList.add("message-header")

  const sender = document.createElement("span")
  sender.classList.add("sender")
  sender.textContent = data.username

  header.appendChild(sender)
  message.appendChild(header)

  const content = document.createElement("div");
  content.classList.add("content")
  content.textContent = data.content

  message.appendChild(content)

  const messages = document.getElementById("messages");
  messages.insertBefore(message, messages.firstChild);

  messages.scrollTop = messages.scrollHeight;
}

const chatDiv = document.getElementById("chat-app");

const ws = new WebSocket("ws://" + location.host + "/ws");

ws.onmessage = (event) => {
  const data = JSON.parse(event.data);

  CreateMessage(data)

};

ws.onclose = () => {
  console.log("Websocket Closed");
};

document
  .getElementById("send-form")
  .addEventListener("htmx:afterRequest", function (evt) {
    const status = evt.detail.xhr.status;

    if (status >= 200 && status < 300) {
      this.reset();
    } else {
      console.warn("Form submission failed with status:", status);
    }
  });
