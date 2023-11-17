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
incrementButton?.addEventListener("click", () => {
  counter.innerText = (parseInt(counter.innerText) + 1).toString();
});
decrementButton?.addEventListener("click", () => {
  counter.innerText = (parseInt(counter.innerText) - 1).toString();
});

// create person
const createPersonButton = document.getElementById(
  "createPersonButton"
) as HTMLButtonElement;
const createdPersonPre = document.getElementById(
  "createdPerson"
) as HTMLPreElement;
createPersonButton?.addEventListener("click", async () => {
  createPersonButton.disabled = true;
  const shariq = await createPerson({ name: "Shariq", id: 1 });
  createPersonButton.disabled = false;
  createdPersonPre.innerText = JSON.stringify(shariq, null, 2);
});

// dashboard modal
document.addEventListener("DOMContentLoaded", () => {
  const dashboardModal = document.getElementById(
    "dashboard-modal"
  ) as HTMLDialogElement;
  if (dashboardModal) {
    const dashboardModalCloseDate = localStorage.getItem(
      "dashboardModalCloseDate"
    );
    if (dashboardModalCloseDate) {
      const modalCloseDate = new Date(dashboardModalCloseDate);
      const today = new Date();
      const diff = today.getTime() - modalCloseDate.getTime();
      const diffInMinutes = diff / (1000 * 60);

      if (diffInMinutes >= 5) {
        (dashboardModal as any).showModal();
      }
    } else {
      (dashboardModal as any).showModal();
    }
  }
});

const dashboardModalCloseButton = document.getElementById(
  "dashbord-modal-close"
) as HTMLButtonElement;
dashboardModalCloseButton?.addEventListener("click", () => {
  localStorage.setItem("dashboardModalCloseDate", new Date().toString());
});

// functions
async function createPerson(person: Person): Promise<Person> {
  return new Promise((resolve) => {
    setTimeout(() => {
      resolve(person);
    }, 2000);
  });
}
