import 'package:jaguar_jwt/jaguar_jwt.dart';
import 'package:sky/config.dart';

final String _secret = Config().jwtAuthSecret;
final String _issuer = 'Fexr Sky';

enum _TokenType {
  access,
  refresh,
}

String _getClaimType(JwtClaim claim) {
  return claim.toJson()['type'];
}

class Token {
  final String token;
  final DateTime expiry;

  Token(this.token, this.expiry);
}

class AccessToken {
  static final int _accessMaxAge = 10; // days
  static get({ required String address, required String f0 }) {
    final String subject = address;

    final JwtClaim claimSet = JwtClaim(
      issuer: _issuer,
      subject: subject,
      maxAge: Duration(days: _accessMaxAge),
      otherClaims: {
        'f0': f0,
        'type': _TokenType.access.toString(),
      }
    );
    return Token(issueJwtHS256(claimSet, _secret), claimSet.expiry!);
  }

  static bool verify(String token) {
    try {
      final JwtClaim claim = verifyJwtHS256Signature(token, _secret);

      claim.validate(
        issuer: _issuer,
      );

      if (_getClaimType(claim) == _TokenType.access.toString()) {
        return true;
      } else {
        return false;
      }
    } on JwtException {
      return false;
    }
  }
}

class RefreshToken {
  static final int _refreshMaxAge = 30; // days

  static Token get({ required String address, required String f0 }) {
    final String subject = address;

    final JwtClaim claimSet = JwtClaim(
      issuer: _issuer,
      subject: subject,
      maxAge: Duration(days: _refreshMaxAge),
      otherClaims: {
        'f0': f0,
        'type': _TokenType.refresh.toString(),
      }
    );
    return Token(issueJwtHS256(claimSet, _secret), claimSet.expiry!);
  }

  static bool verify(String token) {
    try {
      final JwtClaim claim = verifyJwtHS256Signature(token, _secret);

      claim.validate(
        issuer: _issuer,
      );

      if (_getClaimType(claim) == _TokenType.refresh.toString()) {
        return true;
      } else {
        return false;
      }
    } on JwtException {
      return false;
    }
  }
}
