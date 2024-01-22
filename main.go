package main 

import (
	"net/http"
	"errors"
	"github.com/gin-gonic/gin"
)

type sign struct {
	ID		string 	`json:"id"`
	Item	string	`json:"Item"`
	Video	string	`json:"video"`
}

var signssibi = []sign{
	{ID: "SIBI",Item:"A", Video: "A.mp4"},
	{ID: "SIBI",Item:"B", Video: "B.mp4"},
	{ID: "SIBI",Item:"C", Video: "C.mp4"},
	{ID: "SIBI",Item:"D", Video: "D.mp4"},
	{ID: "SIBI",Item:"E", Video: "E.mp4"},
	{ID: "SIBI",Item:"F", Video: "F.mp4"},
	{ID: "SIBI",Item:"G", Video: "G.mp4"},
	{ID: "SIBI",Item:"H", Video: "H.mp4"},
	{ID: "SIBI",Item:"I", Video: "I.mp4"},
	{ID: "SIBI",Item:"J", Video: "J.mp4"},
	{ID: "SIBI",Item:"K", Video: "K.mp4"},
	{ID: "SIBI",Item:"L", Video: "L.mp4"},
	{ID: "SIBI",Item:"M", Video: "M.mp4"},
	{ID: "SIBI",Item:"N", Video: "N.mp4"},
	{ID: "SIBI",Item:"O", Video: "O.mp4"},
	{ID: "SIBI",Item:"P", Video: "P.mp4"},
	{ID: "SIBI",Item:"Q", Video: "Q.mp4"},
	{ID: "SIBI",Item:"R", Video: "R.mp4"},
	{ID: "SIBI",Item:"S", Video: "S.mp4"},
	{ID: "SIBI",Item:"T", Video: "T.mp4"},
	{ID: "SIBI",Item:"U", Video: "U.mp4"},
	{ID: "SIBI",Item:"V", Video: "V.mp4"},
	{ID: "SIBI",Item:"W", Video: "W.mp4"},
	{ID: "SIBI",Item:"X", Video: "X.mp4"},
	{ID: "SIBI",Item:"Y", Video: "Y.mp4"},
	{ID: "SIBI",Item:"Z", Video: "Z.mp4"},
}

var signsbisindo = []sign{
	{ID: "BISINDO",Item:"A", Video: "A.mp4"},
	{ID: "BISINDO",Item:"B", Video: "B.mp4"},
	{ID: "BISINDO",Item:"C", Video: "C.mp4"},
	{ID: "BISINDO",Item:"D", Video: "D.mp4"},
	{ID: "BISINDO",Item:"E", Video: "E.mp4"},
	{ID: "BISINDO",Item:"F", Video: "F.mp4"},
	{ID: "BISINDO",Item:"G", Video: "G.mp4"},
	{ID: "BISINDO",Item:"H", Video: "H.mp4"},
	{ID: "BISINDO",Item:"I", Video: "I.mp4"},
	{ID: "BISINDO",Item:"J", Video: "J.mp4"},
	{ID: "BISINDO",Item:"K", Video: "K.mp4"},
	{ID: "BISINDO",Item:"L", Video: "L.mp4"},
	{ID: "BISINDO",Item:"M", Video: "M.mp4"},
	{ID: "BISINDO",Item:"N", Video: "N.mp4"},
	{ID: "BISINDO",Item:"O", Video: "O.mp4"},
	{ID: "BISINDO",Item:"P", Video: "P.mp4"},
	{ID: "BISINDO",Item:"Q", Video: "Q.mp4"},
	{ID: "BISINDO",Item:"R", Video: "R.mp4"},
	{ID: "BISINDO",Item:"S", Video: "S.mp4"},
	{ID: "BISINDO",Item:"T", Video: "T.mp4"},
	{ID: "BISINDO",Item:"U", Video: "U.mp4"},
	{ID: "BISINDO",Item:"V", Video: "V.mp4"},
	{ID: "BISINDO",Item:"W", Video: "W.mp4"},
	{ID: "BISINDO",Item:"X", Video: "X.mp4"},
	{ID: "BISINDO",Item:"Y", Video: "Y.mp4"},
	{ID: "BISINDO",Item:"Z", Video: "Z.mp4"},
}

func getSigns(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, signssibi)
}

func getSignsBisindo(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, signsbisindo)
}

func addSigns(context *gin.Context) {
	var newSign sign

	if err := context.BindJSON(&newSign); err != nil {
		return
	}

	signssibi = append(signssibi,newSign)

	context.IndentedJSON(http.StatusCreated,newSign)
}

func getSign(context *gin.Context) {
	Item := context.Param("Item")
	s, err := getSignById(Item)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message" : "Sign Not Found"})
		return
	}
	
	context.IndentedJSON(http.StatusOK, s)
}

func getSignBisindo(context *gin.Context) {
	Item := context.Param("Item")
	s, err := getSignBisindoById(Item)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message" : "Sign Not Found"})
		return
	}
	
	context.IndentedJSON(http.StatusOK, s)
}

func getSignById(Item string) (*sign, error) {
	for i, t := range signssibi {
		if t.Item == Item {
			return &signssibi[i], nil
		}
	}

	return nil, errors.New("Sign Not Found")
}

func getSignBisindoById(Item string) (*sign, error) {
	for i, t := range signsbisindo {
		if t.Item == Item {
			return &signsbisindo[i], nil
		}
	}

	return nil, errors.New("Sign Not Found")
}

func main() {
	router := gin.Default()
	router.Static("/video-sibi", "./videos/sibi")
	router.Static("/video-bisindo", "./videos/bisindo")
	router.GET("/sibi-signs", getSigns)
	router.GET("/sibi-signs/:Item", getSign)
	router.GET("/bisindo-signs", getSignsBisindo)
	router.GET("/bisindo-signs/:Item", getSignBisindo)
	router.POST("/addsigns", addSigns)
	router.Run(":8080")
}

