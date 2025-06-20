let ws;
let reconnectAttempts = 0
let reconnecting = false
const messages = document.getElementById("messages");

function CreateMessage(data) {
  const message = document.createElement("div");
  message.classList.add("message");

  const header = document.createElement("div");
  header.classList.add("message-header");

  const sender = document.createElement("span");
  sender.classList.add("sender");
  sender.textContent = data.username;

  header.appendChild(sender);
  message.appendChild(header);

  const content = document.createElement("div");
  content.classList.add("content");
  content.textContent = data.content;

  message.appendChild(content);

  messages.insertBefore(message, messages.firstChild);

  messages.scrollTop = messages.scrollHeight;
}

function connectWebsocket() {
  ws = new WebSocket("ws://" + location.host + "/ws");

  ws.onopen = () => {
    console.log("Websocket opened");
    reconnecting = false

    messages.textContent = ""    
  };

  ws.onmessage = (event) => {
    const data = JSON.parse(event.data);

    CreateMessage(data);
  };

  ws.onerror = (e) => {
    console.error("Websocket error: ", e);
    ws.close();
  };

  ws.onclose = (e) => {
    console.warn("Websocket disconnected: ", e.reason);
    reconnect();
  };
}

function reconnect() {
  if (reconnecting) return
  reconnecting = true

  reconnectAttempts++;
  const timeout = Math.min(1000 * Math.pow(2, reconnectAttempts), 30000);
  console.log(`Reconnecting in ${timeout / 1000}s...`);

  setTimeout(() => {
    connectWebsocket();
  }, timeout);
}

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

connectWebsocket();
