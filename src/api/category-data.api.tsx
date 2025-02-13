import { TreeNode } from "@/components/UI/checkbox-tree";

//返回目录分类数据
export default function CategoryDataApi(): TreeNode[] {
  return [
    {
      id: "1",
      label: "Node 1",
      children: [
        {
          id: "1-1",
          label: "Node 1-1",
        },
        {
          id: "1-2",
          label: "Node 1-2",
          children: [
            {
              id: "1-2-1",
              label: "Node 1-2-1",
            },
          ],
        },
      ],
    },
    {
      id: "2",
      label: "Node 2",
    },
  ];
}