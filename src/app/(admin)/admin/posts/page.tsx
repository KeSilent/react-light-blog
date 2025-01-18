"use client";
import Button from "@/components/UI/button";
import Table from "@/components/UI/table";
import { useRouter } from "next/navigation";

export default function Posts() {
  const router = useRouter();
  const tableThead = ["Post Title", "Author", "Date", "Status"];
  const tableTbody = [
    ["Post Title", "Author", "Date", "Status"],
    ["Post Title", "Author", "Date", "Status"],
    ["Post Title", "Author", "Date", "Status"],
    ["Post Title", "Author", "Date", "Status"],
    ["Post Title", "Author", "Date", "Status"],
    ["Post Title", "Author", "Date", "Status"],
    ["Post Title", "Author", "Date", "Status"],
  ];

  //跳转到新建文章页面
  const handlePostNewClick = () => {
    router.push("/admin/posts/new");
  };

  return (
    <div>
      <div className="pb-4 flex items-center">
        <label className="text-lg font-semibold px-4">文章管理</label>
        <Button title="写文章" onClick={handlePostNewClick} />
      </div>
      <Table tableThead={tableThead} tableTbody={tableTbody}></Table>
    </div>
  );
}
