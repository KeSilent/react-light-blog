import { EchartModel } from "@/model/echart-model";
import { JSX } from "react";

export default function PostAccessSort({ items }: { items: EchartModel[] }) {
  const list: JSX.Element[] = [];
  items.map((item) => {
    list.push(
      <div key={item.name} className="flex justify-between">
        <div>{item.name}</div>
        <div>{item.value}</div>
      </div>
    );
  });

  return (
    <div className="flex flex-col gap-4 w-full">
      <div className="text-2xl font-bold text-secondary border-b-2 m-4 pb-2">
        文章访问Top10
      </div>
      <div className="flex flex-col gap-2 mx-4 text-xl">{list}</div>
    </div>
  );
}
