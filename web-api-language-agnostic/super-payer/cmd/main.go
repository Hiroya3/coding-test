package main

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	pkgErr "super-payer/pkg/error"
	"time"
)

// jwtCustomClaims are custom claims extending default ones.
// See https://github.com/golang-jwt/jwt for more examples
type jwtCustomClaims struct {
	UserID int `json:"user_id"`
	jwt.RegisteredClaims
}

type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type postInvoiceRequest struct {
	ClientID           int       `json:"client_id"`            // 取引先のID
	IssueDate          time.Time `json:"issue_date"`           // UTC時間の発行日
	PayAmount          int       `json:"pay_amount"`           // 請求金額
	Fee                int       `json:"fee"`                  // 手数料
	FeeRate            float64   `json:"fee_rate"`             // 手数料率
	ConsumptionTax     int       `json:"consumption_tax"`      // 消費税
	ConsumptionTaxRate float64   `json:"consumption_tax_rate"` // 消費税率
	PaymentDueDate     string    `json:"payment_due_date"`     // 支払い期日
}

type postInvoiceResponse struct {
	Invoice InvoiceResponse `json:"invoice"`
}

type listInvoiceRequest struct {
	FromDate string `json:"from_date"`
	ToDate   string `json:"to_date"`
}

type ListInvoiceResponse struct {
	Invoices []InvoiceResponse `json:"invoices"`
}

type InvoiceResponse struct {
	InvoiceID          int                        `json:"invoice_id"`
	Company            InvoiceResponseCompany     `json:"company"`
	UserName           string                     `json:"user_name"`
	Client             InvoiceResponseClient      `json:"client"`
	BankAccount        InvoiceResponseBankAccount `json:"bank_account"`
	IssueDate          string                     `json:"issue_date"`
	PayAmount          int                        `json:"pay_amount"`
	Fee                int                        `json:"fee"`
	FeeRate            float64                    `json:"fee_rate"`
	ConsumptionTax     int                        `json:"consumption_tax"`
	ConsumptionTaxRate float64                    `json:"consumption_tax_rate"`
	PaymentDueDate     string                     `json:"payment_due_date"`
	Status             string                     `json:"status"`
}

type InvoiceResponseCompany struct {
	CompanyName        string `json:"company_name"`
	RepresentativeName string `json:"representative_name"`
	PhoneNumber        string `json:"phone_number"`
	PostalCode         string `json:"postal_code"`
	Address            string `json:"address"`
}

type InvoiceResponseClient struct {
	ClientID           string `json:"client_id"`
	CompanyName        string `json:"company_name"`
	RepresentativeName string `json:"representative_name"`
	PhoneNumber        string `json:"phone_number"`
	PostalCode         string `json:"postal_code"`
	Address            string `json:"address"`
}

type InvoiceResponseBankAccount struct {
	BankName      string `json:"bank_name"`
	BranchName    string `json:"branch_name"`
	AccountNumber string `json:"account_number"`
	AccountName   string `json:"account_name"`
}

func login(c echo.Context) error {
	var request loginRequest
	if err := c.Bind(&request); err != nil {
		return err
	}
	username := request.Username
	password := request.Password

	// Throws unauthorized error
	// TODO usersテーブルからハッシュ化されたパスワードの突き合わせをする
	if username != "admin" || password != "password" {
		return echo.ErrUnauthorized
	}

	// Set custom claims
	claims := &jwtCustomClaims{
		1,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	// TODO 秘密鍵を単純でないものにして、envから取得する
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}

// postInvoice : 請求書登録
func postInvoice(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*jwtCustomClaims)
	userID := claims.UserID

	var request postInvoiceRequest
	if err := c.Bind(&request); err != nil {
		return err
	}

	// TODO usecaseと繋ぎこむ
	return convertRes(c, fmt.Sprintf("User ID: %d", userID), nil)
}

// listInvoices : 指定期間内で、userIDが所属する企業の請求書一覧を返す（ページングなし）
func listInvoices(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*jwtCustomClaims)
	userID := claims.UserID

	var request listInvoiceRequest
	if err := c.Bind(&request); err != nil {
		return err
	}

	// TODO usecaseと繋ぎこむ
	return convertRes(c, fmt.Sprintf("User ID: %d", userID), nil)
}

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Login route
	e.POST("/login", login)

	// Restricted group
	r := e.Group("/api/invoices")

	// Configure middleware with the custom claims type
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(jwtCustomClaims)
		},
		SigningKey: []byte("secret"), // TODO 秘密鍵を単純でないものにして、envから取得する
	}
	r.Use(echojwt.WithConfig(config))
	r.GET("", listInvoices)
	r.POST("", postInvoice)

	e.Logger.Fatal(e.Start(":1323"))
}

func convertRes(c echo.Context, res any, err error) error {

	if err == nil {
		// nilであれば正常として返す
		return c.JSON(http.StatusOK, res)
	}

	var originalErr pkgErr.PkgError
	if errors.As(err, &originalErr) {
		switch originalErr.GetKind() {
		case pkgErr.ErrKindNotFound:
			return c.JSON(http.StatusNotFound, res)
		case pkgErr.ErrKindInvalidArgument:
			return c.JSON(http.StatusBadRequest, res)
		case pkgErr.ErrKindInternal:
			return c.JSON(http.StatusInternalServerError, res)
		}
	}

	return c.JSON(http.StatusInternalServerError, res)
}
