# contests

## Description

コンテスト(予選・本戦を管理するため)

<details>
<summary><strong>Table Definition</strong></summary>

```sql
CREATE TABLE `contests` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `title` varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `start_at` timestamp NOT NULL,
  `end_at` timestamp NOT NULL,
  `submit_limit` bigint NOT NULL,
  `slug` varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `tag_selection_logic` enum('auto','manual') COLLATE utf8mb4_bin NOT NULL,
  `validator` varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `time_limit_per_task` bigint DEFAULT '30000000000',
  `created_at` timestamp NOT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `slug` (`slug`),
  UNIQUE KEY `contest_slug` (`slug`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin
```

</details>

## Columns

| Name | Type | Default | Nullable | Extra Definition | Children | Parents | Comment |
| ---- | ---- | ------- | -------- | ---------------- | -------- | ------- | ------- |
| id | bigint |  | false | auto_increment | [submits](submits.md) |  | コンテストID(PK) |
| title | varchar(255) |  | false |  |  |  | コンテスト名 e.g.) 2023 情報科学実験A ベンチマークコンテスト 予選 |
| start_at | timestamp |  | false |  |  |  | 開始日時 |
| end_at | timestamp |  | false |  |  |  | 終了日時 |
| submit_limit | bigint |  | false |  |  |  | グループが提出できる回数(SUCCESS のみカウントされる) |
| slug | varchar(255) |  | false |  |  |  | コンテストの識別子(unique) e.g.) exp-a-2023-qual. |
| tag_selection_logic | enum('auto','manual') |  | false |  |  |  | auto: {slug}/random.txt から選出される / manual: {slug}/\d.txt から選出される(\d は試行回数) |
| validator | varchar(255) |  | false |  |  |  | レスポンスの validator。実装されている validator は proto に定義されている。 |
| time_limit_per_task | bigint | 30000000000 | true |  |  |  | タスクごとの制限時間(ms) |
| created_at | timestamp |  | false |  |  |  | 作成日時 |
| updated_at | timestamp |  | true |  |  |  | 更新日時 |

## Constraints

| Name | Type | Definition |
| ---- | ---- | ---------- |
| contest_slug | UNIQUE | UNIQUE KEY contest_slug (slug) |
| PRIMARY | PRIMARY KEY | PRIMARY KEY (id) |
| slug | UNIQUE | UNIQUE KEY slug (slug) |

## Indexes

| Name | Definition |
| ---- | ---------- |
| PRIMARY | PRIMARY KEY (id) USING BTREE |
| contest_slug | UNIQUE KEY contest_slug (slug) USING BTREE |
| slug | UNIQUE KEY slug (slug) USING BTREE |

## Relations

![er](contests.svg)

---

> Generated by [tbls](https://github.com/k1LoW/tbls)