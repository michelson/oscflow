@Welcome = React.createClass
  render: ->
    <div className="welcome">
      <h1>OscFlow</h1>
      {@.props.children}
      <PatchesList/>
      <MainControl/>
    </div>


@PatchesList = React.createClass
  render: ->
    return(
      <div>
        <p>Patches list</p>
        <ul>
          <li>
            <ReactRouter.Link to="patches/1">Patch 1</ReactRouter.Link>
          </li>
          <li>
            <a href="/patches/2">Patch 2</a>
          </li>
          <li>
            <a href="/patches/3">Patch 3</a>
          </li>
          <li>
            <a href="/patches/4">Patch 4</a>
          </li>
        </ul>
      </div>

    )


@MainControl = React.createClass
  render: ->
    <div>
      <p>Main control</p>
      <ul>
        <li>Play</li>
        <li>Rec</li>
        <li>Volume</li>
      </ul>
    </div>