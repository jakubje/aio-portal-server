package api

import (
	"fmt"
	db "github.com/jakub/aioportal/server/db/sqlc"
	"github.com/jakub/aioportal/server/token"
	"github.com/jakub/aioportal/server/util"

	"github.com/gin-gonic/gin"
)

// Server server HTTP requests for our service
type Server struct {
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
	ctx        gin.Context
}

// NewServer creates a new HTTP server and setup routing
func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}
	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	// user routes
	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)

	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))

	authRoutes.POST("/users/update", server.updateUser)
	authRoutes.POST("/user", server.getUser)
	authRoutes.GET("/users", server.listUsers)
	// related to football so will need to check db migration
	authRoutes.DELETE("/users", server.deleteUser)

	// portfolio routes
	authRoutes.POST("/portfolio", server.createPortfolio)
	authRoutes.POST("/portfolio/update", server.updatePortfolio)
	authRoutes.GET("/portfolio/:id", server.getPortfolio)
	authRoutes.GET("/portfolios/account", server.listPortfolios)
	authRoutes.GET("/portfolio/rollup/:portfolio_id", server.getRollUpByCoinByPortfolio)
	// related to coins - violates foreign key constraint on table coins
	authRoutes.DELETE("/portfolio/:id", server.deletePortfolio)

	// coint routes
	//// need a generic coins database and one for portfolio
	//router.POST("/coin", server.addCoin)
	//router.POST("/coin/update", server.updateCoin)
	//router.GET("/coin/:id", server.getCoin)
	//router.GET("/coins", server.listCoins)
	//router.DELETE("/coin/:id", server.deleteCoin)

	// transaction routes
	authRoutes.POST("/transaction", server.createTransaction)
	authRoutes.GET("/transaction/:id", server.getTransaction)
	authRoutes.GET("/transactions", server.listTransactions)
	authRoutes.GET("/transactions/coin", server.listTransactionsByAccountByCoin)
	authRoutes.DELETE("/transaction/:id", server.deleteTransaction)

	// watchlist routes
	authRoutes.POST("/watchlist", server.createWatchlist)
	authRoutes.POST("/watchlist/update", server.updateWatchlist)
	authRoutes.GET("/watchlist/:id", server.getWatchlist)
	authRoutes.GET("/watchlists", server.listWatchlists)
	authRoutes.DELETE("/watchlist/:id", server.deleteWatchlist)

	// watchlistcoin routes
	authRoutes.POST("/watchlist/coin", server.createWatchlistCoin)
	authRoutes.GET("/watchlist/coin/:id", server.getWatchlistCoin)
	authRoutes.GET("/watchlist/coins/:watchlist_id", server.listWatchlistCoins)
	authRoutes.DELETE("//watchlist/coin/:id", server.deleteWachlistCoin)

	// football routes
	authRoutes.POST("/football", server.addFootball)
	authRoutes.POST("/football/update", server.updateFootball)
	// get returns a EOF error
	authRoutes.GET("/football", server.getFootball)

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
