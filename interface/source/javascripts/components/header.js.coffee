@Header = React.createClass

  render: ->
    return (
      <header className="mdl-layout__header">
        <div className="mdl-layout__header-row">
          {# Title }
          <span className="mdl-layout-title">OscFlow</span>
          {# Add spacer, to align navigation to the right }
          <div className="mdl-layout-spacer"></div>
        </div>
      </header>
    )