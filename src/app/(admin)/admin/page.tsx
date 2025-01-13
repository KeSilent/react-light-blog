import ClassifyTotal from "@/components/admin/dashboard/ClassifyTotal";
import MessageTotal from "@/components/admin/dashboard/MessageTotal";
import PostAccessSort from "@/components/admin/dashboard/PostAccessSort";
import PostTotal from "@/components/admin/dashboard/PostTotal";
import RecentMessages from "@/components/admin/dashboard/RecentMessages";
import RecentPosts from "@/components/admin/dashboard/RecentPosts";
import TypeEChart from "@/components/admin/dashboard/TypeEChart";
import TypeTotal from "@/components/admin/dashboard/TypeTotal";

export default function Home() {
  return (
    <div>
      <div className="flex gap-4 justify-between mt-10">
        <PostTotal total={10} />
        <TypeTotal total={10} />
        <ClassifyTotal total={10} />
        <MessageTotal total={10} />
      </div>
      <div className="flex flex-col md:flex-row justify-between gap-6 shadow-md mt-10">
        <TypeEChart
          chartData={[
            { value: 40, name: "rose 1" },
            { value: 38, name: "rose 2" },
            { value: 32, name: "rose 3" },
            { value: 30, name: "rose 4" },
            { value: 28, name: "rose 5" },
            { value: 26, name: "rose 6" },
            { value: 22, name: "rose 7" },
            { value: 18, name: "rose 8" },
          ]}
        />
        <PostAccessSort
          items={[
            { value: 40, name: "rose 1" },
            { value: 38, name: "rose 2" },
            { value: 32, name: "rose 3" },
            { value: 30, name: "rose 4" },
            { value: 28, name: "rose 5" },
            { value: 26, name: "rose 6" },
            { value: 22, name: "rose 7" },
            { value: 18, name: "rose 8" },
            { value: 18, name: "rose 9" },
            { value: 18, name: "rose 10" },
          ]}
        />
        <TypeEChart
          chartData={[
            { value: 40, name: "rose 1" },
            { value: 38, name: "rose 2" },
            { value: 32, name: "rose 3" },
            { value: 30, name: "rose 4" },
            { value: 28, name: "rose 5" },
            { value: 26, name: "rose 6" },
            { value: 22, name: "rose 7" },
            { value: 18, name: "rose 8" },
          ]}
        />
      </div>
      <div className="flex gap-6 flex-col md:flex-row mt-10">
        <div className="shadow-md w-full">
          <RecentPosts
            tableThead={["标题", "分类", "浏览数", "发布时间"]}
            tableTbody={[
              ["标题111", "分类11", "30", "2025年1月1日"],
              ["标题111", "分类11", "30", "2025年1月1日"],
            ]}
          />
        </div>
        <div className="shadow-md w-full">
          <RecentMessages
            tableThead={["发起人", "内容", "出处", "评论时间"]}
            tableTbody={[
              ["标题111", "分类11", "30", "2025年1月1日"],
              ["标题111", "分类11", "30", "2025年1月1日"],
            ]}
          />
        </div>
      </div>
    </div>
  );
}
