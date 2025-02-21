export default function Table({
  tableThead,
  tableTbody,
}: {
  tableThead: string[];
  tableTbody: string[][];
}) {
  const thead = tableThead.map((item, index) => {
    return (
      <th key={index} className="p-2">
        {item}
      </th>
    );
  });
  const tbody = tableTbody.map((item, index) => {
    return (
      <tr key={index}>
        {item.map((item, index) => {
          return (
            <td key={index} className="p-2">
              {item}
            </td>
          );
        })}
      </tr>
    );
  });

  return (
    <table className="table-auto w-full">
      <thead className="text-left border-b-2 text-secondary">
        <tr>{thead}</tr>
      </thead>
      <tbody>{tbody}</tbody>
    </table>
  );
}
