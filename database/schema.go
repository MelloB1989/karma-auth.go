package database

import "time"

type Organisation struct {
    OID       string    `json:"oid"`
    Name      string    `json:"name"`
    Email     string    `json:"email"`
    CreatedAt time.Time `json:"createdAt"`
    Details   string    `json:"details"`
    Contact   string    `json:"contact"`
    Domain    string    `json:"domain"`
}

type UpdateOrganisation struct {
    OID     string  `json:"oid"`
    Name    *string `json:"name,omitempty"`
    Email   *string `json:"email,omitempty"`
    Details *string `json:"details,omitempty"`
    Contact *string `json:"contact,omitempty"`
    Domain  *string `json:"domain,omitempty"`
}

type ApiToken struct {
    TID       string    `json:"tid"`
    OID       string    `json:"oid"`
    Token     string    `json:"token"`
    CreatedAt time.Time `json:"createdAt"`
    ExpiresAt time.Time `json:"expiresAt"`
    Details   string    `json:"details"`
}

type UpdateApiToken struct {
    TID       string     `json:"tid"`
    Token     *string    `json:"token,omitempty"`
    ExpiresAt *time.Time `json:"expiresAt,omitempty"`
    Details   *string    `json:"details,omitempty"`
}

type User struct {
    UID                    string    `json:"uid"`
    OID                    string    `json:"oid"`
    Email                  string    `json:"email"`
    Phone                  string    `json:"phone"`
    Name                   string    `json:"name"`
    CreatedAt              time.Time `json:"createdAt"`
    Details                string    `json:"details"`
    Password               string    `json:"password"`
    LastPassword           string    `json:"lastPassword"`
    PasswordResetToken     string    `json:"passwordResetToken"`
    PasswordResetExpires   time.Time `json:"passwordResetExpires"`
    EmailVerified          bool      `json:"emailVerified"`
    EmailVerificationToken string    `json:"emailVerificationToken"`
    PhoneVerified          bool      `json:"phoneVerified"`
}

type UpdateUser struct {
    Email                  *string    `json:"email,omitempty"`
    Phone                  *string    `json:"phone,omitempty"`
    Name                   *string    `json:"name,omitempty"`
    Details                *string    `json:"details,omitempty"`
    Password               *string    `json:"password,omitempty"`
    LastPassword           *string    `json:"lastPassword,omitempty"`
    PasswordResetToken     *string    `json:"passwordResetToken,omitempty"`
    PasswordResetExpires   *time.Time `json:"passwordResetExpires,omitempty"`
    EmailVerified          *bool      `json:"emailVerified,omitempty"`
    EmailVerificationToken *string    `json:"emailVerificationToken,omitempty"`
    PhoneVerified          *bool      `json:"phoneVerified,omitempty"`
}

type UpdateUserDetails struct {
    Email   *string `json:"email,omitempty"`
    Phone   *string `json:"phone,omitempty"`
    Name    *string `json:"name,omitempty"`
    Details *string `json:"details,omitempty"`
}

type Role struct {
    RID       string    `json:"rid"`
    OID       string    `json:"oid"`
    Name      string    `json:"name"`
    CreatedAt time.Time `json:"createdAt"`
    Allows    string    `json:"allows"`
    Denies    string    `json:"denies"`
}

type UpdateRole struct {
    Name   *string `json:"name,omitempty"`
    Allows *string `json:"allows,omitempty"`
    Denies *string `json:"denies,omitempty"`
}

type UserRole struct {
    UID       string    `json:"uid"`
    RID       string    `json:"rid"`
    CreatedAt time.Time `json:"createdAt"`
}

type Permissions struct {
    Allows []string `json:"allows"`
    Denies []string `json:"denies"`
}