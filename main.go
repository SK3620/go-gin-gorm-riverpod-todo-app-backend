package main

import (
	"go-gin-gorm-riverpod-todo-app/controllers"
	"go-gin-gorm-riverpod-todo-app/infra"
	"go-gin-gorm-riverpod-todo-app/middlwares"

	// "go-gin-gorm-riverpod-todo-app/models"
	"go-gin-gorm-riverpod-todo-app/repositories"
	"go-gin-gorm-riverpod-todo-app/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

/*
3層アーキテクチャ構成
① Controller → リクエストデータのハンドリングやレスポンスの設定
② Service → 実現したい機能の実装（ビジネスロジック）
③ Repository → データの永続化やデータソースとのやりとり（メモリ上 or DB）

// 依存性の流れ
Controller → Service → Repositorie → Database

Controller → Service
コントローラはリクエストを処理するが、ビジネスロジックはサービスに委ねる。依存性を注入することで、コントローラはサービスの具体的な実装に依存しない。

Service → Repository
サービスはビジネスロジックを管理するが、（メモリ or DB上の）データ操作はリポジトリに委ねる。依存性を注入することで、サービスはリポジトリの具体的な実装に依存しない。

Repository → Database
リポジトリはデータベース操作を管理する。外部（main関数）で生成したDBのインスタンスを依存性として注入する。
*/

func setupRouter(db *gorm.DB) *gin.Engine {
	// メモリ上(NewTodoMemoryRepository)でデータ操作を行う場合、以下のサンプルデータを使用する
	/*
	// サンプルデータを作成
	todos := []models.Todo{
		{ID: 1, Title: "タイトル１", IsCompleted: false},
		{ID: 2, Title: "タイトル２", IsCompleted: true},
		{ID: 3, Title: "タイトル３", IsCompleted: false},
	}
	*/

	// ファクトリ関数を用いてそれぞれ依存性を注入していく
	
	// 以下todoRespositoryを片方に切り替えることでControllerとServiceを修正することなく、容易にデータソースのやり取り先を変更できる
	// todoRepository := repositories.NewTodoMemoryRepository(todos) // →　メモリ上でデータを操作
	todoRepository := repositories.NewTodoRepository(db) //　→ DB上でデータを操作
	todoService := services.NewTodoService(todoRepository)
	todoController := controllers.NewTodoController(todoService)

	authRepository := repositories.NewAuthRepository(db)
	authService := services.NewAuthService(authRepository)
	authController := controllers.NewAuthController(authService)

	r := gin.Default()

	// 共通のルートをグループ化
	authRouter := r.Group("/auth")
	//　AuthControllerでリクエストデータをハンドリングする前に、認証ミドルウェアを挟み、JWTToken認証を実施
	todoRouterWithAuth := r.Group("/todos", middlwares.AuthMiddlware(authService))

	authRouter.POST("/sign_up", authController.SignUp) // サインアップ
	authRouter.POST("/login", authController.Login) // ログイン

	todoRouterWithAuth.GET("", todoController.FindAll) // ログインしたユーザーに紐ずくtodoを全件取得
	todoRouterWithAuth.POST("", todoController.Create) // todo新規作成
	todoRouterWithAuth.PUT("/:id", todoController.Update) // todoの更新
	todoRouterWithAuth.DELETE("/:id", todoController.Delete) // todoの削除

	return r
}

// エントリーポイント
func main() {
	// envファイルを読み込む
	infra.Initialize()

	// DBへの接続
	db := infra.SetupDB()

	// APIルートを定義
	r := setupRouter(db)
	
	r.Run("localhost:8080")
}
