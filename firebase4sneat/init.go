package firebase4sneat

import (
	"github.com/sneat-co/sneat-core-modules/auth/token4auth"
	"github.com/sneat-co/sneat-go-server/facade2firebase"
	"github.com/sneat-co/sneat-go-server/sneatfb"
)

func InitFirebase() {
	firebaseProjectID := sneatfb.GetFirebaseProjectID()
	sneatfb.InitFirebaseForSneat(firebaseProjectID, "sneat")
	token4auth.IssueAuthToken = facade2firebase.IssueFirebaseAuthToken

	// TODO: Investigate why the code below has been commented out
	//facade4auth.CreateUserInAuth = facade2firebase.CreateFirebaseUser
	//facade4auth.DeleteUserFromAuth = facade2firebase.DeleteFirebaseUser
}
