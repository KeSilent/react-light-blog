import Table from "@/components/UI/table";

export default function RecentPosts({
  tableThead,
  tableTbody,
}: {
  tableThead: string[];
  tableTbody: string[][];
}) {
  return <Table tableThead={tableThead} tableTbody={tableTbody} />;
}
