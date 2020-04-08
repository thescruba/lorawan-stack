// Copyright © 2019 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import React from 'react'
import PropTypes from 'prop-types'
import { Map, Marker, TileLayer } from 'react-leaflet'
import classnames from 'classnames'
import Leaflet from 'leaflet'
import MarkerIcon from '../../assets/auxiliary-icons/location_pin.svg'

import style from './map.styl'

// delete Leaflet.Icon.Default.prototype._getIconUrl
Leaflet.Icon.Default.mergeOptions({
  iconRetinaUrl: MarkerIcon,
  iconUrl: MarkerIcon,
  iconSize: [26, 36],
  shadowSize: [26, 36],
  iconAnchor: [13, 36],
  shadowAnchor: [8, 37],
  popupAnchor: [0, -40],
  shadowUrl: require('leaflet/dist/images/marker-shadow.png'),
})

export default class LocationMap extends React.Component {
  static propTypes = {
    // LeafletConfig is an object which can contain any number of properties defined by the leaflet plugin and is used to overwrite the default configuration of leaflet.
    leafletConfig: PropTypes.shape({}),
    // Markers is an array of objects containing a specific properties
    markers: PropTypes.arrayOf(
      // Position is a object containing two properties latitude and longitude which are both numbers.
      PropTypes.shape({
        position: PropTypes.shape({
          longitude: PropTypes.number,
          latitude: PropTypes.number,
        }),
      }),
    ),
    onClick: PropTypes.func,
    // Widget is a boolean used to add a class name to the map container div for styling.
    widget: PropTypes.bool,
  }

  static defaultProps = {
    leafletConfig: {},
    widget: false,
    markers: [],
    onClick: undefined,
  }

  render() {
    const { widget, markers, onClick, leafletConfig } = this.props
    const position =
      markers.length >= 1
        ? [markers[0].position.latitude, markers[0].position.longitude]
        : [52.7, 4.9]
    const map = (
      <Map
        className={classnames(style.map, { [style.widget]: widget })}
        center={position}
        zoom={7}
        minZoom={7}
        onClick={onClick}
        {...leafletConfig}
      >
        <TileLayer
          url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
          attribution='&copy; <a href="http://osm.org/copyright">OpenStreetMap</a> contributors'
        />
        {markers.map(
          (marker, idx) =>
            marker && (
              <Marker
                key={`marker-${idx}`}
                position={[marker.position.latitude, marker.position.longitude]}
              />
            ),
        )}
      </Map>
    )
    return <div className={style.container}>{map}</div>
  }
}
