import Table from "@/components/UI/table";

export default function RecentMessages({
  tableThead,
  tableTbody,
}: {
  tableThead: string[];
  tableTbody: string[][];
}) {
  return <Table tableThead={tableThead} tableTbody={tableTbody} />;
}
