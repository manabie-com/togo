describe('Create todo test case', () => {
    it('verify limit 5 task per user per day flow', () => {
        for (let i = 0; i < 6; i++) {
            cy.request({
                method: 'POST',
                url: "http://localhost:3001/createTodo",
                body: {
                    "content": "asdasd1123sdzasdasdvcs23123",
                    "status": "ACTIVE",
                    "user_id": "asdasdasd" // test case only pass one time, need to change user to check others
                },
                failOnStatusCode: false,
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': 'testabc.xyz.ahk'
                }
            }).then((res) => {
                if (i == 5) {
                    expect(res.status).to.eq(400)
                    expect(res.body.error).to.eq("Limit of 5 tasks asdasdasd can be added per day.")
                } else {
                    expect(res.status).to.eq(200)
                }
            })

        }

    })
})
