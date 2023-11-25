const button = document.getElementById("dashboard-modal-button");
button?.addEventListener("click", async () => {
  const { camelCase } = await import("lodash");
  console.log(camelCase("hello world"));
});
