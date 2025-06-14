const ws = new WebSocket("ws://" + location.host + "/ws")

ws.onmessage = (event) => {
    const data = event.data

    const div = document.createElement("div")
    div.classList.add("message")
    div.textContent = data

    const messages = document.getElementById("messages")
    messages.appendChild(div)

    messages.scrollTop = messages.scrollHeight
}

ws.onclose = () => {
    console.log("Websocket Closed")
}

document.getElementById('send-form').addEventListener('htmx:afterRequest', function(evt) {
    this.reset();
});
