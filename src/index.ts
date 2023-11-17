interface Person {
  name: string;
  id: number;
}

// increment state
const incrementButton = document.getElementById(
  "incrementButton"
) as HTMLButtonElement;
const decrementButton = document.getElementById(
  "decrementButton"
) as HTMLButtonElement;
const counter = document.getElementById("counter") as HTMLDivElement;
incrementButton.addEventListener("click", () => {
  counter.innerText = (parseInt(counter.innerText) + 1).toString();
});
decrementButton.addEventListener("click", () => {
  counter.innerText = (parseInt(counter.innerText) - 1).toString();
});

// create person
const createPersonButton = document.getElementById(
  "createPersonButton"
) as HTMLButtonElement;
const createdPersonPre = document.getElementById(
  "createdPerson"
) as HTMLPreElement;
createPersonButton.addEventListener("click", async () => {
  createPersonButton.disabled = true;
  const shariq = await createPerson({ name: "Shariq", id: 1 });
  createPersonButton.disabled = false;
  createdPersonPre.innerText = JSON.stringify(shariq, null, 2);
});

// functions
async function createPerson(person: Person): Promise<Person> {
  return new Promise((resolve) => {
    setTimeout(() => {
      resolve(person);
    }, 2000);
  });
}
