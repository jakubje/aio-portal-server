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
	router.POST("/users/login", server.loginUser)
	router.POST("/users", server.createUser)
	router.POST("/users/update/:id", server.updateUser)
	router.POST("/user", server.getUser)
	router.GET("/users", server.listUsers)
	// related to football so will need to check db migration
	router.DELETE("/users/:id", server.deleteUser)

	// portfolio routes
	router.POST("/portfolio", server.createPortfolio)
	router.POST("/portfolio/update", server.updatePortfolio)
	router.GET("/portfolio/:id", server.getPortfolio)
	router.GET("/portfolio/account/:account_id", server.listPortfolios)
	router.GET("/portfolio/rollup/:portfolio_id", server.getRollUpByCoinByPortfolio)
	// related to coins - violates foreign key constraint on table coins
	router.DELETE("/portfolio/:id", server.deletePortfolio)

	// coint routes
	//// need a generic coins database and one for portfolio
	//router.POST("/coin", server.addCoin)
	//router.POST("/coin/update", server.updateCoin)
	//router.GET("/coin/:id", server.getCoin)
	//router.GET("/coins", server.listCoins)
	//router.DELETE("/coin/:id", server.deleteCoin)

	// transaction routes
	router.POST("/transaction", server.createTransaction)
	router.GET("/transaction/:id", server.getTransaction)
	router.GET("/transactions", server.listTransactions)
	router.GET("/transactions/coin", server.listTransactionsByAccountByCoin)
	router.DELETE("/transaction/:id", server.deleteTransaction)

	// watchlist routes
	router.POST("/watchlist", server.createWatchlist)
	router.POST("/watchlist/update", server.updateWatchlist)
	router.GET("/watchlist/:id", server.getWatchlist)
	router.GET("/watchlists/:account_id", server.listWatchlists)
	router.DELETE("/watchlist/:id", server.deleteWatchlist)

	// watchlistcoin routes
	router.POST("/watchlist/coin", server.createWatchlistCoin)
	router.GET("/watchlist/coin/:id", server.getWatchlistCoin)
	router.GET("/watchlist/coins/:watchlist_id", server.listWatchlistCoins)
	router.DELETE("//watchlist/coin/:id", server.deleteWachlistCoin)

	// football routes
	router.POST("/football", server.addFootball)
	router.POST("/football/update", server.updateFootball)
	// get returns a EOF error
	router.GET("/football/:id", server.getFootball)

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
