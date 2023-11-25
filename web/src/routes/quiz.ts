type Variant = "multiple-choice" | "open-ended";

const multipleChoiceButton = document.getElementById(
  "multiple-choice-button"
) as HTMLButtonElement;
const openEndedButton = document.getElementById(
  "open-ended-button"
) as HTMLButtonElement;
const variantInput = document.getElementById("variant-input");

// switching variant
multipleChoiceButton?.addEventListener("click", () => {
  if (variantInput?.getAttribute("value") !== ("multiple-choice" as Variant)) {
    openEndedButton.classList.remove("btn-neutral");
    multipleChoiceButton.classList.add("btn-neutral");
    variantInput?.setAttribute("value", "multiple-choice" as Variant);
  }
});
openEndedButton?.addEventListener("click", () => {
  if (variantInput?.getAttribute("value") !== ("open-ended" as Variant)) {
    multipleChoiceButton.classList.remove("btn-neutral");
    openEndedButton.classList.add("btn-neutral");
    variantInput?.setAttribute("value", "open-ended" as Variant);
  }
});
