package templates

const RCE = `import React, { Component } from 'react'

export class rce extends Component {
    render() {
        return (
            <div>

            </div>
        )
    }
}

export default rce`
const RFCE = `import React from 'react'

function rfce() {
    return (
        <div>

        </div>
    )
}

export default rfce`
