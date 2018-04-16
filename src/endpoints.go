package main

const (
	// プロデューサーの作成
	// 一度アクセスするとuuidが生成されるので、以降はPRODUCER_INFORMATIONから操作していく
	CREATE_PRODUCER = "/createProducer"

	// プロデューサー情報の確認(GET)・追加と変更(POST)
	PRODUCER_INFORMATION = "/producerContents"

	// アイドルの作成
	// 一度アクセスするとuuidが生成されるので、以降はIDOL_CONTENTSから操作していく
	CREATE_IDOL = "/createIdol"

	// アイドルコンテンツの確認(GET)・追加(POST)
	IDOL_CONTENTS = "/idolContents"


	// ローカルに存在するProducerとIdolの情報をjsonで返す
	// 主にロード画面で使う
	PRODUCER_LIST = "/producerList"
	IDOL_LIST = "/idolList"

	// IDOL_LISTの返り値にあるUUIDをクエリにつけてプロデュース開始
	// ユーザーがクライアントを落とした場合、回収するためにフラグをここで立てる
	START_PRODUCE = "/startProduce"


	// 招待イベントや活動週、流行の情報、週間ハイライトをまとめて返すラッパー
	MORNING_CHECK = "/morningCheck"
	// 中身
	WEEK_CHECK = "/weekCheck" //もし1週だった場合、上がらなければ引退になる
	TREND_CHECK = "/trendCheck"
	INVITE_CHECK = "/inviteCheck"
	WEEK_HIGHLIGHT = "/weekHighlight"

	// SCHEDULEにGETでスケジュール確認、POSTで予定変更ができる
	SCHEDULE = "/schedule"
	// 中身
	IDOL_SCHEDULE = "/idolSchedule"
	PRODUCER_SCHEDULE = "/producerSchedule"

	// ACTIVITY_LISTでスケジュールに登録できる行動を確認する(GET)
	ACTIVITY_LIST = "/activityList"
	// 中身
	IDOL_ACTIVITY_LIST = "/idolActivityList"
	PRODUCER_ACTIVITY_LIST = "/producerActivityList"


	// スケジュールに行動を登録しようとすると以下にアクセスし、更に詳細な選択肢を渡す
	// 例:レッスンを選択 > アクセスが発生して選択肢表示 > ボーカルレッスンを選択 という流れ
	AUDITION_LIST = "/auditionList"
	LESSON_LIST = "/lessonList"
	EVENT_LIST = "/eventList"

	NEGOTIATION_LIST = "/negotiationList" //仕事交渉
	VISIT_LIST = "/visitList" //あいさつ回り

	COMMU_LIST = "/commuList"
	LIVE_LIST = "/liveList"


	// SCHEDULEにある情報から活動を始める
	START_WEEK = "/startWeek"

	// カットインイベントを含め、ほとんどのゲーム内行動はここで帰ってくる。
	// レスポンスに合わせてUIの操作を行う(コミュの選択肢あたりは手元で処理)
	ACTIVITY = "/activity"


	// コミュの結果をポストする
	// 手元では選択の結果で文章が変わるだけで、好感度やフラグなどはポストされてから鯖側で処理する
	COMMU_RESULT = "/commuResult"


	// レポート類を受け取る
	NIGHT_REPORT = "/nightReport"
	// 中身
	IDOL_ACTIVITY_REPORT = "/idolActivityReport"
	FAN_REPORT = "/fanReport"

	// リミットの確認
	// もしリミットなら次の確認が飛び、LAST_CONCERTかZ_ENDのどちらかになる
	LIMIT_CHECK = "/limitCheck"

	// 自主活動の確認
	// GETしてからPOSTにするか、組み込みにしてPOSTするか考え中
	PLAN_CHECK = "/planCheck"


	// プロデュースを一旦休む
	// START_PRODUCER時に立てたフラグを回収して正常終了とする
	REST_PRODUCE = "/restProduce"



	// ======以下は不必要でありながら作りたいと思うエンドポイント====== //
	// =====他の実装を進めつつ、本当に実装したいか自分と考えてみる===== //

	LAST_CONCERT = "/lastConcert"
	Z_END = "/zEnd"

	// 最終処理
	// アクセスする必要性はないんだけど、最後のアクセスとして作りたかった。
	END_PRODUCE = "/endProduce"
)