@Patch = React.createClass

  render: ->
    console.log(@props.params.patch)
    return (
      <div>
        <h3>This the patch {@props.params.patch}</h3>
      </div>
    ) 