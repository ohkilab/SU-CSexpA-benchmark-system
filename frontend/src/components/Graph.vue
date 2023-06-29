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

  const startDate = new Date('2023-06-27')

  const numberOfDays = Math.ceil((Date.now() - startDate) / (1000 * 60 * 60 * 24)) + 1

  const chartData = {
    labels: Array.from(new Array(numberOfDays)).map((el, i) => {
      startDate.setDate(startDate.getDate() + 1)

      return startDate.toLocaleDateString()
    }),
    datasets: [
      {
        label: 'A01',
        backgroundColor: '#f87979',
        data: [400, 500, 500, 600]
      },
      {
        label: 'A02',
        backgroundColor: '#777',
        data: [80, 400, 1021, 1232]
      },
      {
        label: 'A03',
        backgroundColor: '#777',
        data: [200, 900, 900, 1023]
      }
    ]
  }

  const chartOptions = {
    responsive: true,
    maintainAspectRatio: false
  }
</script>
<template>
    <Line
        id="id"
        :options="chartOptions"
        :data="chartData"
    />
</template>
