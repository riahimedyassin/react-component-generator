export const RCE = `import React, { Component } from 'react'

export class REPLACE_NAME extends Component {
    render() {
        return (
            <div>

            </div>
        )
    }
}

export default REPLACE_NAME`;

export const RFCE = `import React from 'react'

function REPLACE_NAME() {
    return (
        <div>

        </div>
    )
}

export default REPLACE_NAME`;
