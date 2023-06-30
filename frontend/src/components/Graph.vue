<script setup lang="ts">
import {
  Chart as ChartJS,
  CategoryScale,
  LinearScale,
  PointElement,
  LineElement,
  Title,
  Tooltip,
  Legend
} from 'chart.js'
import { useStateStore } from '../stores/state'
import { Line } from 'vue-chartjs'
import { onMounted } from 'vue'
import { Status, Submit } from 'proto-gen-web/services/backend/resources'

ChartJS.register(
  CategoryScale,
  LinearScale,
  PointElement,
  LineElement,
  Title,
  Tooltip,
  Legend
)

const state = useStateStore()

const random = (min: number, max: number) => Math.floor(Math.random() * (max - min) + min)

const formatDate = (timestamp: number): string => {
  const dateObject: Date = new Date(timestamp * 1000)
  const date: string = dateObject.toLocaleDateString()
  const time: string = dateObject.toLocaleTimeString()

  return `${date}`
}

const today = new Date()
const startDate = new Date()
// one day before actual start date
//const startDate = new Date('2023-06-26')

// show 4 days before
startDate.setDate(today.getDate() - 5)


const numberOfDays = Math.floor((Date.now() - Number(startDate)) / (1000 * 60 * 60 * 24))

/** schema
    [{date: '',scores: [{group: '', score: 0}]}]
*/

// TODO: move fetchSubmits to state actions

interface GraphData {
  date: string,
  maxScores: {
    group: string,
    score: number
  }[]
}

const colors:any = {
  'A1': '#331832',
  'A2': '#D81E5B',
  'A3': '#F0544F',
  'A4': '#C6D8D3',
  'A5': '#FDF0D5',
  'A6': '#DDC4DD',
  'A7': '#DCCFEC',
  'A8': '#4F517D'
}

const graphData:GraphData[] = Array.from(new Array(numberOfDays)).map((el, i) => {
  startDate.setDate(startDate.getDate() + 1)

  return {
    date: startDate.toLocaleDateString(),
    maxScores: state.submits // filter by Status.SUCCESS -> filter by date -> sort by desc score -> separate by group (should return highest score)
      .filter(submit => submit.status === Status.SUCCESS)
      .filter(submit => startDate.toLocaleDateString() === formatDate(Number(submit.submitedAt?.seconds)))
      .sort((a, b) => b.score - a.score)
      .reduce((acc:any, cur:Submit) => { // separate scores by group
        if(cur.groupName in acc) {
          return acc
        } else {
          acc[cur.groupName] = cur
        }
        return acc
      }, {})
  }
})

if(import.meta.env.DEV) console.log('graphdata',graphData)

const chartData = {
  labels: graphData.map(g => g.date),
  datasets: state.records.map(record => {
    return {
      label: record.group?.id,
      backgroundColor: colors[record.group?.id ?? 0] ?? '#f0f', // unity material error color if color not found
      // data: [random(0, 400), random(500, 600), random(600, 800), random(900, 4000)]
      data: graphData.map((g: any) => g.maxScores[record.group?.id ?? 0]?.score ?? null)
    }
  })
}
const chartOptions = {
  responsive: true,
  maintainAspectRatio: false
}

onMounted(() => {
  // if(import.meta.env.DEV) BigInt.prototype.toJSON = function() {return this.toString()}
  if(import.meta.env.DEV) console.log('chartdata', chartData)
})
</script>
<template>
  <Line id="id" :options="chartOptions" :data="chartData" />
</template>
