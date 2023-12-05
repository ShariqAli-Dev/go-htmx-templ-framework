// dashboard modal
const dashboardModal = document.getElementById(
  "dashboard-modal"
) as HTMLDialogElement;
const dashboardModalCloseButton = document.getElementById(
  "dashbord-modal-close"
) as HTMLButtonElement;
const dashboardModalOpenButton = document.getElementById(
  "open-dashboard-modal-button"
) as HTMLButtonElement;

document.addEventListener("DOMContentLoaded", () => {
  if (dashboardModal) {
    const dashboardModalCloseDate = localStorage.getItem(
      "dashboardModalCloseDate"
    );
    if (dashboardModalCloseDate) {
      const modalCloseDate = new Date(dashboardModalCloseDate);
      const today = new Date();
      const diff = today.getTime() - modalCloseDate.getTime();
      const diffInDays = diff / (1000 * 60 * 60 * 24);

      if (diffInDays >= 1) {
        dashboardModal.showModal();
      }
    } else {
      dashboardModal.showModal();
    }
  }
});

dashboardModalCloseButton?.addEventListener("click", () => {
  localStorage.setItem("dashboardModalCloseDate", new Date().toString());
});

dashboardModalOpenButton?.addEventListener("click", async () => {
  dashboardModal.showModal();
});

// word map
const wordmapInput = document.getElementById(
  "wordmap-words"
) as HTMLInputElement;
const margin = { x: 20, y: 20 };
const width = 450 - 2 * margin.x;
const height = 450 - 2 * margin.y;
const words = (JSON.parse(wordmapInput.value) as string[]).map((word) => {
  return { text: word, size: 10 + Math.random() * 90 };
});
var svg = d3
  .select("#word-cloud")
  .append("svg")
  .attr("width", width + margin.x)
  .attr("height", height + margin.y)
  .append("g")
  .attr("transform", "translate(" + margin.x / 2 + "," + margin.y / 2 + ")");
const layout = d3.layout
  .cloud()
  .size([width, height])
  .words(words)
  .padding(10)
  .fontSize(40)
  .on("end", draw);
layout.start();

function draw(words: any) {
  svg
    .append("g")
    .attr(
      "transform",
      "translate(" + layout.size()[0] / 2 + "," + layout.size()[1] / 2 + ")"
    )
    .selectAll("text")
    .data(words)
    .enter()
    .append("text")
    .style("font-size", function (d: any) {
      return d.size + "px";
    })
    .attr("text-anchor", "middle")
    .attr("transform", function (d: any) {
      return "translate(" + [d.x, d.y] + ")rotate(" + d.rotate + ")";
    })
    .text(function (d: any) {
      return d.text;
    })
    .on("click", function (word: any, _) {
      window.location.href =
        window.location.origin + "/quiz" + "?topic=" + (word as any).text;
    })
    .style("cursor", "pointer");
}
