const createGame = document.getElementById("createGame");
const dialog = document.getElementById("dialog");

createGame.onclick = () => {
    dialog.showModal();
};

document.addEventListener("DOMContentLoaded", () => {
    console.log("HELLO");
});
