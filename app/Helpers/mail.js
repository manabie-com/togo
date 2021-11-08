import sgMail from '@sendgrid/mail'
import fs from 'fs'
export function sendMail(params) {
    if(!params.subject){
        return "Không để trống tiêu đề"
    }
    if(!params.to){
        return "Không để trống người nhận mail"
    }
    if(!params.content && !params.html){
        return "Nội dung mail không để trống"
    }
    const msg = {
        to: params.to, // Change to your recipient
        from: params.from, // Change to your verified sender
        subject: params.subject,
        text: params.content,
        html: params.html,
        attachments: []
    }
    let content
    if(params.attachments){
        for(let a of params.attachments){
            let ext = a.split(".").pop().toLowerCase()
            switch (ext) {
                case "pdf":
                    content = fs.readFileSync(a).toString("base64")
                    msg.attachments.push({
                        content: content,
                        filename: a.split("\\").pop(), //lay ten file
                        type: "application/pdf",
                        disposition: "attachment"
                    })
                    break
                case "zip":
                    content = fs.readFileSync(a).toString("base64")
                    msg.attachments.push({
                        content: content,
                        filename: a.split("\\").pop(), //lay ten file
                        type: "application/zip",
                        disposition: "attachment"
                    })
                    break

            }
        }
    }
    sgMail.send(msg)
}