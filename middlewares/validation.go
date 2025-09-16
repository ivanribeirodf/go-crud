package middlewares

import (
    "net/http"
    "strings"

    "github.com/gin-gonic/gin"
    "github.com/go-playground/validator/v10"
)

func ValidationErrorHandler(c *gin.Context) {
    c.Next() // executa a request e coleta os erros depois

    if len(c.Errors) > 0 {
        var errors []string
        for _, e := range c.Errors {
            // Se for erro de validação
            if ve, ok := e.Err.(validator.ValidationErrors); ok {
                for _, fe := range ve {
                    field := fe.Field()
                    tag := fe.Tag()
                    var msg string

                    switch tag {
                    case "required":
                        msg = field + " é obrigatório"
                    case "email":
                        msg = field + " deve ser um email válido"
                    case "min":
                        msg = field + " deve ter pelo menos " + fe.Param() + " caracteres"
                    default:
                        msg = field + " inválido"
                    }

                    errors = append(errors, msg)
                }
            } else {
                errors = append(errors, e.Error())
            }
        }

        c.JSON(http.StatusBadRequest, gin.H{
            "errors": strings.Join(errors, ", "),
        })
        c.Abort()
    }
}
