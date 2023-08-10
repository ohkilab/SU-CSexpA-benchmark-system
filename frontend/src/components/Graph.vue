<script setup lang="ts">
import {
  Chart as ChartJS,
  CategoryScale,
  LinearScale,
  PointElement,
  LineElement,
  Title,
  Tooltip,
  Legend,
} from "chart.js";
import { useStateStore } from "../stores/state";
import { Line } from "vue-chartjs";
import { onMounted } from "vue";
import { Status, Submit } from "proto-gen-web/services/backend/resources";

ChartJS.register(
  CategoryScale,
  LinearScale,
  PointElement,
  LineElement,
  Title,
  Tooltip,
  Legend,
);

const state = useStateStore();

const formatDate = (timestamp: number): string => {
  const dateObject: Date = new Date(timestamp * 1000);
  const date: string = dateObject.toLocaleDateString();

  return `${date}`;
};

const today = new Date();
const startDate = new Date();
// one day before actual start date
//const startDate = new Date('2023-06-26')

// show 7 days before
startDate.setDate(today.getDate() - 8);

const numberOfDays = Math.floor(
  (Date.now() - Number(startDate)) / (1000 * 60 * 60 * 24),
);

/** schema
    [{date: '',scores: [{group: '', score: 0}]}]
*/

// TODO: move fetchSubmits to state actions

interface GraphData {
  date: string;
  maxScores: {
    group: string;
    score: number;
  }[];
}

// color pallete from https://lospec.com/palette-list/city28
const colors: any = {
  A1: "#ceddf0",
  A2: "#8a9ee6",
  A3: "#7169d1",
  A4: "#664db3",
  A5: "#4f2e87",
  A6: "#3a1a59",
  A7: "#45054d",
  A8: "#6b328c",
  A9: "#84449c",
  A10: "#8e60b3",
  A11: "#8f82c2",
  A12: "#f5b771",
  A13: "#e8865f",
  A14: "#f75e5e",
  A15: "#f03554",
  B1: "#b8185d",
  B2: "#730858",
  B3: "#f2bda2",
  B4: "#d6988b",
  B5: "#c77d77",
  B6: "#a65b5b",
  B7: "#8a3f54",
  B8: "#782a49",
  B9: "#521036",
  B10: "#85de8b",
  B11: "#41a681",
  B12: "#207d75",
  B13: "#104957",
};

const graphData: GraphData[] = Array.from(new Array(numberOfDays)).map((_) => {
  startDate.setDate(startDate.getDate() + 1);

  return {
    date: startDate.toLocaleDateString(),
    maxScores: state.submits // filter by Status.SUCCESS -> filter by date -> sort by desc score -> separate by group (should return highest score)
      .filter((submit) => submit.status === Status.SUCCESS)
      .filter(
        (submit) =>
          startDate.toLocaleDateString() ===
          formatDate(Number(submit.submitedAt?.seconds)),
      )
      .sort((a, b) => b.score - a.score)
      .reduce((acc: any, cur: Submit) => {
        // separate scores by group
        if (cur.groupName in acc) {
          return acc;
        } else {
          acc[cur.groupName] = cur;
        }
        return acc;
      }, {}),
  };
});

if (import.meta.env.DEV) console.log("graphdata", graphData);

const chartData = {
  labels: graphData.map((g) => g.date),
  datasets: state.records
    .filter((record) => (record.score ?? 0) > 0)
    .map((record) => {
      return {
        label: record.group?.name,
        backgroundColor: colors[record.group?.name ?? 0] ?? "#f0f", // unity material error color if color not found
        // data: [random(0, 400), random(500, 600), random(600, 800), random(900, 4000)]
        data: graphData.map(
          (g: any) => g.maxScores[record.group?.name ?? 0]?.score ?? null,
        ),
      };
    }),
};
const chartOptions = {
  responsive: true,
  maintainAspectRatio: false,
};

onMounted(() => {
  // if(import.meta.env.DEV) BigInt.prototype.toJSON = function() {return this.toString()}
  if (import.meta.env.DEV) console.log("chartdata", chartData);
});
</script>
<template>
  <Line id="id" :options="chartOptions" :data="chartData" />
</template>
