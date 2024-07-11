package mypage

import (
	"log/slog"
	"net/http"

	"oblivion/pkg/model"

	"oblivion/pkg/crud"
	"oblivion/pkg/discord"
	"time"

	"oblivion/pkg/session"

	"github.com/gin-gonic/gin"
)

func getUserInfo(token string) error {
	tmUserinfo, err := discord.GetUserInfo(token)
	if err != nil {
		return err
	}

	disUserInfo = *tmUserinfo
	return nil
}

var disUserInfo discord.UserInfoResponse

func MypageTop() gin.HandlerFunc {
	return func(c *gin.Context) {
		data := session.Default(c, "session", &model.Session_model{}).Get(c)

		// セッション情報が存在しない場合はログイン画面にリダイレクト
		if data == nil || data.(*model.Session_model).User.UserId == "" {
			c.Redirect(http.StatusSeeOther, "/auth/login")
			return
		}

		if data.(*model.Session_model).Token != "" {
			err := getUserInfo(data.(*model.Session_model).Token)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			// 存在しなければユーザ情報をDBに登録
			userid, err := crud.GetUserId(disUserInfo.Username)
			if err != nil {
				err = crud.InsertUser(disUserInfo.Username, disUserInfo.Email, "")
				if err != nil {
					slog.Error(err.Error())
					c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
					return
				}
			}

			se_data := model.Session_model{
				SessionId: "",
				CookieKey: "session",
				User: model.User{
					UserId:       userid,
					UserName:     disUserInfo.Username,
					Comportement: model.Comportement{Id: "CP-" + userid},
				},
				Token: data.(*model.Session_model).Token,
			}
			// セッション情報を更新
			session.Default(c, "session", &model.Session_model{}).Set(c, se_data)
		}

		// 24時間以内にリマインドがある要素を取得する
		elements, err := crud.GetListElement(data.(*model.Session_model).User.UserId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		doElements := []model.Element{}
		expiredElement := []model.Element{}
		for _, element := range elements {
			// リマインド日時が24時間以内の要素を取得する
			remindTime, err := time.ParseInLocation("2006-01-02 15:04:05", element.Remind, time.Local)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			// 現在の日付を取得
			today := time.Now().Format("2006-01-02")

			// remindTimeが今日の日付と同じであれば追加, 期限切れであればexpiredElementに追加
			if remindTime.Format("2006-01-02") == today {
				doElements = append(doElements, element)
			} else if remindTime.Format("2006-01-02") < today {
				expiredElement = append(expiredElement, element)
			}
		}

		c.HTML(http.StatusOK, "mypage.html", gin.H{
			"user":            data.(*model.Session_model).User,
			"elements":        doElements,
			"expiredElements": expiredElement,
			"IsAuthenticated": true,
		})
	}
}
