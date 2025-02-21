"use client";
import * as echarts from "echarts/core";
import { PieChart } from "echarts/charts";
import { LegendComponent, ToolboxComponent } from "echarts/components";
import { CanvasRenderer } from "echarts/renderers";
import React, { useEffect, useRef } from "react";
import { LabelLayout } from "echarts/features";
import { EchartModel } from "@/model/echart-model";

// 注册必须的组件
echarts.use([
  ToolboxComponent,
  LegendComponent,
  PieChart,
  CanvasRenderer,
  LabelLayout,
]);

export default function TypeEChart({
  chartData,
}: {
  chartData: EchartModel[];
}) {
  const chartRef = useRef<HTMLDivElement>(null);

  useEffect(() => {
    if (!chartRef.current) return;

    const myChart = echarts.init(chartRef.current);
    const option = {
      legend: {
        top: "bottom",
      },
      toolbox: {
        show: true,
        feature: {
          mark: { show: true },
          dataView: { show: true, readOnly: false },
          restore: { show: true },
          saveAsImage: { show: true },
        },
      },
      series: [
        {
          name: "文章分类统计",
          type: "pie",
          radius: [50, 150],
          center: ["50%", "50%"],
          roseType: "area",
          itemStyle: {
            borderRadius: 8,
          },
          data: chartData,
        },
      ],
    };
    myChart.setOption(option);

    return () => {
      myChart.dispose();
    };
  }, [chartData]);

  return (
    <div className="flex flex-col w-full">
      <div className="text-2xl font-bold text-secondary border-b-2 m-4 pb-2">爬虫访问统计Top10</div>
      <div>
        <div ref={chartRef} style={{ width: "100%", height: "400px" }}></div>
      </div>
    </div>
  );
}
