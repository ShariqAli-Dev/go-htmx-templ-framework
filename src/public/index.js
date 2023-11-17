// increment state
const incrementButton = document.getElementById("incrementButton");
const decrementButton = document.getElementById("decrementButton");
const counter = document.getElementById("counter");
incrementButton.addEventListener("click", () => {
    counter.innerText = (parseInt(counter.innerText) + 1).toString();
});
decrementButton.addEventListener("click", () => {
    counter.innerText = (parseInt(counter.innerText) - 1).toString();
});
// create person
const createPersonButton = document.getElementById("createPersonButton");
const createdPersonPre = document.getElementById("createdPerson");
createPersonButton.addEventListener("click", async () => {
    createPersonButton.disabled = true;
    const shariq = await createPerson({ name: "Shariq", id: 1 });
    createPersonButton.disabled = false;
    createdPersonPre.innerText = JSON.stringify(shariq, null, 2);
});
// functions
async function createPerson(person) {
    return new Promise((resolve) => {
        setTimeout(() => {
            resolve(person);
        }, 2000);
    });
}
