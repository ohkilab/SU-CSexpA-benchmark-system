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
import { Status } from 'proto-gen-web/services/backend/resources'

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

  const formatDate = (timestamp: number):string => {
    const dateObject: Date = new Date(timestamp * 1000)
    const date: string = dateObject.toLocaleDateString()
    const time: string = dateObject.toLocaleTimeString()
    
    return `${date}`
  }

  // one day before actual start date
  const startDate = new Date('2023-06-26')

  const numberOfDays = Math.floor((Date.now() - Number(startDate)) / (1000 * 60 * 60 * 24))

  const graphData = Array.from(new Array(numberOfDays)).map((el, i) => {
    startDate.setDate(startDate.getDate() + 1)

    return {
      date: startDate.toLocaleDateString(),
      maxScore: state.submits // filter by Status.SUCCESS -> filter by date -> sort by desc score -> get first index
                  .filter(submit => submit.status === Status.SUCCESS)
                  .filter(submit => startDate.toLocaleDateString() === formatDate(Number(submit.submitedAt?.seconds)))
                  .sort((a, b) => b.score - a.score)
                  .at(0)
                  ?.score
    }
  })

  const chartData = {
    // labels: Array.from(new Array(numberOfDays)).map((el, i) => {
    //   // startDate.setDate(startDate.getDate() + 1)
    //
    //   return startDate.toLocaleDateString()
    // }),
    labels: graphData.map(g => g.date),
      // {
      //   label: 'A01',
      //   backgroundColor: '#f87979',
      //   data: [400, 500, 500, 600]
      // },
    datasets: state.records.map(record => {
      return {
        label: record.group?.id,
        backgroundColor: `#${record.rank}${record.rank}${record.rank}`,
        // data: [random(0, 400), random(500, 600), random(600, 800), random(900, 4000)]
        data: graphData.map(g => g.maxScore)
      }
    })
  }
  const chartOptions = {
    responsive: true,
    maintainAspectRatio: false
  }

  onMounted(() => {
    // if(import.meta.env.DEV) BigInt.prototype.toJSON = function() {return this.toString()}
  })
</script>
<template>
    <!-- <pre>{{state.submits.map(submit => formatDate(Number(submit.submitedAt?.seconds)))}}</pre> -->
    <!-- <pre>{{graphData}}</pre> -->
    <Line
        id="id"
        :options="chartOptions"
        :data="chartData"
    />
</template>
