package firebase4sneat

import (
	"github.com/sneat-co/sneat-core-modules/auth/token4auth"
	"github.com/sneat-co/sneat-go-firebase/facade2firebase"
	"github.com/sneat-co/sneat-go-firebase/sneatfb"
)

func InitFirebase() {
	firebaseProjectID := sneatfb.GetFirebaseProjectID()
	sneatfb.InitFirebaseForSneat(firebaseProjectID, "sneat")
	//
	//facade4auth.CreateUserInAuth = facade2firebase.CreateFirebaseUser
	//facade4auth.DeleteUserFromAuth = facade2firebase.DeleteFirebaseUser
	token4auth.IssueAuthToken = facade2firebase.IssueFirebaseAuthToken
}
