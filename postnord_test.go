package postnord

import (
	"encoding/xml"
	"testing"
)

var t1 = `
<?xml version="1.0" encoding="UTF-8" ?>
<TrackingInformationResponse>
	<shipments>
		<Shipment>
			<shipmentId>RG136027285CN</shipmentId>
			<uri>/ntt-service-rest/api/shipment/RG136027285CN/0</uri>
			<assessedNumberOfItems>1</assessedNumberOfItems>
			<service>
				<code>RG</code>
				<name>Rekommenderat brev</name>
			</service>
			<consignee>
				<address/>
			</consignee>
			<statusText>
				<header>Sändningen är utlämnad till mottagaren</header>
			</statusText>
			<status>DELIVERED</status>
			<items>
				<Item>
					<itemId>RG136027285CN</itemId>
					<deliveryDate>2015-03-03T17:10:00</deliveryDate>
					<noItems>1</noItems>
					<status>DELIVERED</status>
					<statusText>
						<header>Försändelsen är utlämnad till mottagaren</header>
						<body>Försändelsen lämnades ut 2015-03-03 kl 17:10</body>
					</statusText>
					<events>
						<TrackingEvent>
							<eventTime>2015-02-26T17:08:00</eventTime>
							<eventCode>VAPO_70</eventCode>
							<eventDescription>Avsändaren har lämnat in försändelsen i utlandet</eventDescription>
							<location>
								<locationId>CN</locationId>
								<displayName>Kina</displayName>
								<name>Kina</name>
								<locationType>UNDEF</locationType>
							</location>
						</TrackingEvent>
						<TrackingEvent>
							<eventTime>2015-03-02T07:31:00</eventTime>
							<eventCode>VAPO_68</eventCode>
							<eventDescription>Försändelsen har kommit från avsändarlandet till Postens utrikesterminal för sortering</eventDescription>
							<location>
								<locationId>SESTOA</locationId>
								<displayName>Stockholm utr,Sverige</displayName>
								<name>Stockholm utr,Sverige</name>
								<locationType>UNDEF</locationType>
							</location>
						</TrackingEvent>
						<TrackingEvent>
							<eventTime>2015-03-03T07:54:00</eventTime>
							<eventCode>VAPO_114</eventCode>
							<eventDescription>Försändelsen är vidaresänd till mottagarens utlämningsställe</eventDescription>
							<location>
								<locationId>354531</locationId>
								<displayName>Posten Nyköping</displayName>
								<name>Posten Nyköping</name>
								<postcode>61138</postcode>
								<locationType>SERVICE_POINT</locationType>
							</location>
						</TrackingEvent>
						<TrackingEvent>
							<eventTime>2015-03-03T10:46:00</eventTime>
							<eventCode>VAPO_30</eventCode>
							<eventDescription>Försändelsen har kommit till mottagarens utlämningsställe. Express körs ut till mottagaren, övriga försändelser aviseras</eventDescription>
							<location>
								<locationId>354665</locationId>
								<displayName>Willys Nyköping</displayName>
								<name>Willys Nyköping</name>
								<postcode>61138</postcode>
								<locationType>SERVICE_POINT</locationType>
							</location>
						</TrackingEvent>
						<TrackingEvent>
							<eventTime reference="../../../deliveryDate"/>
							<eventCode>VAPO_28</eventCode>
							<eventDescription>Försändelsen är överlämnad till mottagaren</eventDescription>
							<location>
								<locationId>354665</locationId>
								<displayName>Willys Nyköping</displayName>
								<name>Willys Nyköping</name>
								<postcode>61138</postcode>
								<locationType>SERVICE_POINT</locationType>
							</location>
						</TrackingEvent>
					</events>
					<references/>
					<itemRefIds/>
					<freeTexts/>
				</Item>
			</items>
			<additionalServices/>
			<splitStatuses/>
			<shipmentReferences/>
		</Shipment>
	</shipments>
</TrackingInformationResponse>
`

func Test_xmlDecode(t *testing.T) {
	var s shipmentResponse
	err := xml.Unmarshal([]byte(t1), &s)
	if err != nil {
		t.Error(err)
	}

	t.Logf("%+v", s)

	if s.Shipments[0].ShipmentId != "RG136027285CN" {
		t.Fail()
	}

	if len(s.Shipments) < 1 {
		t.Fail()
	}
}
