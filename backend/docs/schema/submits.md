# submits

## Description

グループの提出を管理するテーブル

<details>
<summary><strong>Table Definition</strong></summary>

```sql
CREATE TABLE `submits` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `url` varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `score` bigint DEFAULT NULL,
  `language` enum('php','go','rust','javascript','csharp','cpp','ruby','python') COLLATE utf8mb4_bin DEFAULT NULL,
  `message` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL,
  `status` varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `task_num` bigint NOT NULL,
  `submited_at` timestamp NOT NULL,
  `completed_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `contest_submits` bigint DEFAULT NULL,
  `group_submits` bigint DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `submits_contests_submits` (`contest_submits`),
  KEY `submits_groups_submits` (`group_submits`),
  CONSTRAINT `submits_contests_submits` FOREIGN KEY (`contest_submits`) REFERENCES `contests` (`id`) ON DELETE SET NULL,
  CONSTRAINT `submits_groups_submits` FOREIGN KEY (`group_submits`) REFERENCES `groups` (`id`) ON DELETE SET NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin
```

</details>

## Columns

| Name | Type | Default | Nullable | Extra Definition | Children | Parents | Comment |
| ---- | ---- | ------- | -------- | ---------------- | -------- | ------- | ------- |
| id | bigint |  | false | auto_increment | [task_results](task_results.md) |  | 提出ID(PK) |
| url | varchar(255) |  | false |  |  |  | 提出の URL |
| score | bigint |  | true |  |  |  | task_results の合計スコア |
| language | enum('php','go','rust','javascript','csharp','cpp','ruby','python') |  | true |  |  |  |  |
| message | varchar(255) |  | true |  |  |  | connection_failed や validation_error などのメッセージ |
| status | varchar(255) |  | false |  |  |  | ステータス。proto に定義されている。 |
| task_num | bigint |  | false |  |  |  | タスクの数 |
| submited_at | timestamp |  | false |  |  |  | 提出日時 |
| completed_at | timestamp |  | true |  |  |  | ベンチマーク完了日時 |
| updated_at | timestamp |  | true |  |  |  |  |
| contest_submits | bigint |  | true |  |  | [contests](contests.md) |  |
| group_submits | bigint |  | true |  |  | [groups](groups.md) |  |

## Constraints

| Name | Type | Definition |
| ---- | ---- | ---------- |
| PRIMARY | PRIMARY KEY | PRIMARY KEY (id) |
| submits_contests_submits | FOREIGN KEY | FOREIGN KEY (contest_submits) REFERENCES contests (id) |
| submits_groups_submits | FOREIGN KEY | FOREIGN KEY (group_submits) REFERENCES groups (id) |

## Indexes

| Name | Definition |
| ---- | ---------- |
| submits_contests_submits | KEY submits_contests_submits (contest_submits) USING BTREE |
| submits_groups_submits | KEY submits_groups_submits (group_submits) USING BTREE |
| PRIMARY | PRIMARY KEY (id) USING BTREE |

## Relations

![er](submits.svg)

---

> Generated by [tbls](https://github.com/k1LoW/tbls)