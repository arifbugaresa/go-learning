-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE email_templates (
   id               SERIAL PRIMARY KEY,
   code             VARCHAR(255) NOT NULL UNIQUE,
   name             VARCHAR(255) NOT NULL,
   template         TEXT NOT NULL,
   created_at       TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
   created_by       VARCHAR(356) DEFAULT 'SYSTEM',
   modified_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
   modified_by      VARCHAR(356) DEFAULT 'SYSTEM'
);

INSERT INTO email_templates (code, name, template)
VALUES ('login_email_template', 'Login Email Template', '<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><meta name="viewport" content="width=device-width,initial-scale=1"><title>Notification Email</title></head><body style="font-family:Arial,sans-serif;background-color:#f4f4f4;margin:0;padding:0"><div style="max-width:600px;margin:20px auto;background-color:#fff;border:1px solid #e0e0e0;border-radius:5px;box-shadow:0 4px 8px rgba(0,0,0,.1)"><div style="background-color:#007bff;color:#fff;padding:20px;text-align:center;border-top-left-radius:5px;border-top-right-radius:5px"><h1 style="margin:0;font-size:24px">Notification</h1></div><div style="padding:20px;font-size:16px;line-height:1.6;color:#333"><h2 style="color:#007bff">Hello, {{.Name}}</h2><p>We wanted to let you know about a recent login activity on your account.</p><table style="width:100%;border-collapse:collapse;margin-top:10px"><tr><th style="text-align:left;padding:8px;background-color:#f4f4f4;border:1px solid #ddd;color:#007bff">Device</th><td style="padding:8px;border:1px solid #ddd">{{.Device}}</td></tr><tr><th style="text-align:left;padding:8px;background-color:#f4f4f4;border:1px solid #ddd;color:#007bff">IP Address</th><td style="padding:8px;border:1px solid #ddd">{{.IPAddress}}<br></td></tr><tr><th style="text-align:left;padding:8px;background-color:#f4f4f4;border:1px solid #ddd;color:#007bff">Location</th><td style="padding:8px;border:1px solid #ddd">{{.Location}}</td></tr><tr><th style="text-align:left;padding:8px;background-color:#f4f4f4;border:1px solid #ddd;color:#007bff">Time</th><td style="padding:8px;border:1px solid #ddd">{{.LoginTime}}</td></tr></table><p>If you do not recognize this login activity or believe there might be an issue with your account,<a href="mailto:support@example.com" style="color:#007bff;text-decoration:none"> contact our support team</a> immediately to ensure your account''s security.</p></div><div style="background-color:#f4f4f4;padding:10px;text-align:center;font-size:12px;color:#666;border-bottom-left-radius:5px;border-bottom-right-radius:5px"><p>Leaaning Project 2024 &copy; All rights reserved.</p></div></div></body></html>');

-- +migrate StatementEnd